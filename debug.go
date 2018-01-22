package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"

	"github.com/ccding/go-stun/stun"
	ptp "github.com/subutai-io/p2p/lib"
)

type DaemonArgs struct {
	IP         string `json:"ip"`
	Mac        string `json:"mac"`
	Dev        string `json:"dev"`
	Hash       string `json:"hash"`
	Dht        string `json:"dht"`
	Keyfile    string `json:"keyfile"`
	Key        string `json:"key"`
	TTL        string `json:"ttl"`
	Fwd        bool   `json:"fwd"`
	Port       int    `json:"port"`
	Interfaces bool   `json:"interfaces"` // show only
	All        bool   `json:"all"`        // show only
	Command    string `json:"command"`
	Args       string `json:"args"`
	Log        string `json:"log"`
}

// Debug prints debug information
func CommandDebug(restPort int) {
	out, err := sendRequest(restPort, "debug", &DaemonArgs{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(out.Message)
	os.Exit(out.Code)
}

// ExecDaemon starts P2P daemon
func ExecDaemon(port int, sFile, profiling, syslog string) {
	if syslog != "" {
		ptp.SetSyslogSocket(syslog)
	}
	StartProfiling(profiling)
	go ptp.InitPlatform()
	ptp.InitErrors()

	if !ptp.CheckPermissions() {
		os.Exit(1)
	}

	ReadyToServe = false
	proc := new(Daemon)
	proc.Initialize(sFile)
	setupRESTHandlers(port, proc)

	ptp.Log(ptp.Info, "Determining outbound IP")
	nat, host, err := stun.NewClient().Discover()
	if err != nil {
		ptp.Log(ptp.Error, "Failed to discover outbound IP: %s", err)
		OutboundIP = nil
	} else {
		OutboundIP = net.ParseIP(host.IP())
		ptp.Log(ptp.Info, "Public IP is %s. %s", OutboundIP.String(), nat)
	}

	if sFile != "" {
		ptp.Log(ptp.Info, "Restore file provided")
		// Try to restore from provided file
		instances, err := proc.Instances.LoadInstances(proc.SaveFile)
		if err != nil {
			ptp.Log(ptp.Error, "Failed to load instances: %v", err)
		} else {
			ptp.Log(ptp.Info, "%d instances were loaded from file", len(instances))
			for _, inst := range instances {
				proc.run(&inst, new(Response))
			}
		}
	}

	ReadyToServe = true

	SignalChannel = make(chan os.Signal, 1)
	signal.Notify(SignalChannel, os.Interrupt)

	go func() {
		for sig := range SignalChannel {
			fmt.Println("Received signal: ", sig)
			pprof.StopCPUProfile()
			os.Exit(0)
		}
	}()
	select {}
}

func (d *Daemon) execRESTDebug(w http.ResponseWriter, r *http.Request) {
	if !ReadyToServe {
		resp, _ := getResponse(105, "P2P Daemon is in initialization state")
		w.Write(resp)
	}
	args := new(DaemonArgs)
	err := getJSON(r.Body, args)
	if handleMarshalError(err, w) != nil {
		return
	}
	response := new(Response)
	d.Debug(&Args{
		Command: args.Command,
		Args:    args.Args,
	}, response)
	resp, err := getResponse(response.ExitCode, response.Output)
	if err != nil {
		ptp.Log(ptp.Error, "Internal error: %s", err)
		return
	}
	w.Write(resp)
}

// Debug output debug information
func (p *Daemon) Debug(args *Args, resp *Response) error {
	ptp.Log(ptp.Info, "Preparing Debug output")
	resp.Output = "DEBUG INFO:\n"
	resp.Output = fmt.Sprintf("Version: %s Build: %s\n", AppVersion, BuildID)
	resp.Output += fmt.Sprintf("Number of gouroutines: %d\n", runtime.NumGoroutine())
	resp.Output += fmt.Sprintf("Instances information:\n")
	instances := p.Instances.Get()
	for _, inst := range instances {
		resp.Output += fmt.Sprintf("Bootstrap nodes:\n")
		for _, conn := range inst.PTP.Dht.Connections {
			resp.Output += fmt.Sprintf("\t%s\n", conn.RemoteAddr().String())
		}
		resp.Output += fmt.Sprintf("Hash: %s\n", inst.ID)
		resp.Output += fmt.Sprintf("ID: %s\n", inst.PTP.Dht.ID)
		resp.Output += fmt.Sprintf("UDP Port: %d\n", inst.PTP.UDPSocket.GetPort())
		resp.Output += fmt.Sprintf("Interface %s, HW Addr: %s, IP: %s\n", inst.PTP.Interface.GetName(), inst.PTP.Interface.GetHardwareAddress().String(), inst.PTP.Interface.GetIP().String())
		resp.Output += fmt.Sprintf("Proxies:\n")
		if len(inst.PTP.Proxies) == 0 {
			resp.Output += fmt.Sprintf("\tNo proxies in use\n")
		}
		for _, proxy := range inst.PTP.Proxies {
			resp.Output += fmt.Sprintf("\tProxy address: %s\n", proxy.Addr.String())
		}
		resp.Output += fmt.Sprintf("Peers:\n")

		peers := inst.PTP.Peers.Get()
		for _, peer := range peers {
			resp.Output += fmt.Sprintf("\t--- %s ---\n", peer.ID)
			if peer.PeerLocalIP == nil {
				resp.Output += "\t\tNo IP assigned\n"

			} else if peer.PeerHW == nil {
				resp.Output += "\t\tNo MAC assigned\n"
			} else {
				resp.Output += fmt.Sprintf("\t\tHWAddr: %s\n", peer.PeerHW.String())
				resp.Output += fmt.Sprintf("\t\tIP: %s\n", peer.PeerLocalIP.String())
				resp.Output += fmt.Sprintf("\t\tEndpoint: %s\n", peer.Endpoint)
				resp.Output += fmt.Sprintf("\t\tPeer Address: %s\n", peer.PeerAddr.String())
				proxyInUse := "No"
				if peer.IsUsingTURN {
					proxyInUse = "Yes"
				}
				resp.Output += fmt.Sprintf("\t\tUsing proxy: %s\n", proxyInUse)
			}
			resp.Output += fmt.Sprintf("\t--- End of %s ---\n", peer.ID)
		}
	}
	return nil
}
