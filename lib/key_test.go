package ptp

import (
	"reflect"
	"testing"
	"time"
)

func TestKey_generateID(t *testing.T) {
	type fields struct {
		id      string
		key     string
		added   time.Time
		expires time.Time
	}
	type args struct {
		idList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Generating key ID", fields{key: "some random crypto key"}, args{}, false},
		{"Generating key ID: Failing case", fields{key: "some random crypto key"}, args{[]string{"24c58e0", "a14b742", "a856a08",
			"195c1f7", "70d6f8e", "9fae00b", "fca56d6", "040bc1d", "e4a0587", "e05bff2", "4a06036", "4136e69", "f1ae3e3", "4930c68",
			"bed0c1d", "4a2751b", "d930743", "6478e6a", "eab14dd", "a775fae", "6aaa4b3", "259e7ad", "f3756f2", "fa968ab", "1718fc6",
			"aa88fa2", "0d2eb8d", "f196de1", "05b7834", "e7e02dc", "1c9265c", "ac38718", "2daf262", "76f906f", "9f4f075", "11f28fe",
			"09787eb", "7be70a9", "5aac692", "54ec72e", "b43c91e", "208ddb3", "f2b89b2", "ac4dd54", "867f455", "f0070b1", "5a2166d",
			"07b5863", "4e7a32d", "cb57392", "97eeae1", "37ae1d7", "f28e793", "a05eda8", "b67679c", "f74f390", "e847eca", "1025400",
			"9e72221", "c889e85", "d1f8d61", "cb2fedd", "5782d89", "7d16d15", "ba5ef84", "783cabb", "cd7d080", "e20851b", "eb63c3c",
			"2fe8316", "2ef0ae0", "247bfb4", "2cdb9a0", "64ed00a", "8143bdd", "657ad66", "39fe37c", "cab6ae7", "b60d61e", "b472b43",
			"16088f4", "0bbaaf8", "712417a", "4cacd86", "13e4ea1", "5eefdc6", "52b251b", "31617e6", "911ceac", "bc9ce04", "849c110",
			"8e67968", "b6f0dd9", "f908b50", "f8bb6fc", "6ad1112", "c64d245", "4369f43", "6417332", "348711b", "d6786fc", "c08660e",
			"bfba517", "86118cd", "c1fa135", "6a50e72", "78b1f67", "6ddd1fe", "7de6953", "5ceee41", "4baf2fc", "1a7099e", "fe94e2a",
			"01bcc4b", "511bb48", "cf29839", "bdc7ef2", "c7c464b", "a852bbf", "9446cef", "04aad2a", "3b5d49e", "0916eb1", "edbb6b3",
			"9080154", "5fd2b44", "4e770d5", "0e0e9f1", "be3fa94", "3a6e3b5", "8095f9a", "3d96ad5", "0f6a77f", "c1a9fc9", "0198bca",
			"9e335b5", "f37aae0", "675f92d", "df302c3", "fc8fce2", "7c01740", "1d783a4", "94c022d", "29ae28a", "d2f98d8", "b6b7afd",
			"2d2f428", "4e5dd4a", "99ace42", "e6c598a", "2c15500", "5855b33", "fbaa78e", "a1a6b39", "693bbe8", "93bb503", "24af328",
			"302f89b", "daec02b", "29e41a5", "727a3e5", "7806a09", "be2bd98", "ffa2fb7", "2b8664d", "d918764", "9953034", "511173a",
			"6aa0220", "8b6e268", "cad2e5e", "4fa6cb7", "dc638bc", "e80deda", "e8b54ce", "d4ccf50", "9600f4c", "61c42c3", "7d9bc1c",
			"3bf6ad6", "0cff02c", "6a9eff0", "18ed7c9", "b753d3d", "0d1f3ec", "8d5913f", "5e8dca6", "a42c076", "34e4569", "6aa6770",
			"448775d", "8198459", "7d09ddd", "7490362", "3590c41", "a53d4aa", "b3e3222", "1531d01", "74a25ce", "e8a9e64", "1ea60c2",
			"f8a5388", "0bd7cc0", "086fde3", "1d3ea58", "49a1583", "6981bdd", "127723c", "c5d29dd", "b955f68", "c66a555", "003b462",
			"975a4fd", "1c50c79", "4224494", "9e17c6a", "17a2adc", "067c65a", "f31c114", "134ae3f", "0ae89c8", "f793b11", "1422acc",
			"d5b29f4", "b6a05e2", "566fd9a", "19a4872", "cb7cb40", "86ff510", "b2d8807", "455e1fe", "3d11d9c", "883320e", "b6c7f0d",
			"b5fe379", "30fb1bb", "3fe0042", "5ec7928", "791b7a9", "f7be1d1", "43ef0ff", "97bafd3", "a0e367b", "ca8044c", "04d5f1e",
			"9594a35", "fef578e", "96e93c4", "fe2ba58", "9198a2b", "ffc29c4", "63bf2eb", "23aea2b", "7915e5c", "6181720", "dcb1cc4"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				id:      tt.fields.id,
				key:     tt.fields.key,
				added:   tt.fields.added,
				expires: tt.fields.expires,
			}
			if err := k.generateID(tt.args.idList); (err != nil) != tt.wantErr {
				t.Errorf("Key.generateID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKey_IsExpired(t *testing.T) {
	type fields struct {
		id      string
		key     string
		added   time.Time
		expires time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Not expired key", fields{expires: time.Now().Local().Add(time.Hour * time.Duration(1))}, false},
		{"Expired key", fields{expires: time.Unix(1, 1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				id:      tt.fields.id,
				key:     tt.fields.key,
				added:   tt.fields.added,
				expires: tt.fields.expires,
			}
			if got := k.IsExpired(); got != tt.want {
				t.Errorf("Key.IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKey(t *testing.T) {
	type args struct {
		key        string
		start      time.Time
		expiration time.Time
		idList     []string
	}
	key := new(Key)
	key.expires = time.Now().Local().Add(time.Hour * time.Duration(1))
	key.key = "working_key"
	key.generateID([]string{})
	tests := []struct {
		name    string
		args    args
		want    *Key
		wantErr bool
	}{
		{"Expired key", args{expiration: time.Unix(1, 1)}, nil, true},
		{"Passing case", args{expiration: key.expires, key: "working_key"}, key, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKey(tt.args.key, tt.args.start, tt.args.expiration, tt.args.idList)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if !reflect.DeepEqual(got.key, tt.want.key) {
					t.Errorf("NewKey() = %v, want %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewKey() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestKey_IsStarted(t *testing.T) {
	type fields struct {
		id      string
		key     string
		added   time.Time
		starts  time.Time
		expires time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Not started key", fields{starts: time.Now().Local().Add(time.Hour * time.Duration(1))}, false},
		{"Started key", fields{starts: time.Unix(1, 1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				id:      tt.fields.id,
				key:     tt.fields.key,
				added:   tt.fields.added,
				starts:  tt.fields.starts,
				expires: tt.fields.expires,
			}
			if got := k.IsStarted(); got != tt.want {
				t.Errorf("Key.IsStarted() = %v, want %v", got, tt.want)
			}
		})
	}
}
