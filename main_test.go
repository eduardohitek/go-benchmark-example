package main

import "testing"

func Test_palindromeReverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Banana", args: args{str: "banana"}, want: false},
		{name: "Civic", args: args{str: "civic"}, want: true},
		{name: "Level", args: args{str: "level"}, want: true},
		{name: "Noon", args: args{str: "noon"}, want: true},
		{name: "Eduardo", args: args{str: "eduardo"}, want: false},
		{name: "Rodador", args: args{str: "rodador"}, want: true},
		{name: "Rir", args: args{str: "rir"}, want: true},
		{name: "Ovo", args: args{str: "ovo"}, want: true},
		{name: "15651", args: args{str: "15651"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromeReverse(tt.args.str); got != tt.want {
				t.Errorf("palindromeReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindromeFromEnd(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Banana", args: args{str: "banana"}, want: false},
		{name: "Civic", args: args{str: "civic"}, want: true},
		{name: "Level", args: args{str: "level"}, want: true},
		{name: "Noon", args: args{str: "noon"}, want: true},
		{name: "Eduardo", args: args{str: "eduardo"}, want: false},
		{name: "Rodador", args: args{str: "rodador"}, want: true},
		{name: "Rir", args: args{str: "rir"}, want: true},
		{name: "Ovo", args: args{str: "ovo"}, want: true},
		{name: "15651", args: args{str: "15651"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromeFromEnd(tt.args.str); got != tt.want {
				t.Errorf("palindromeFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindromeFromMiddle(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Banana", args: args{str: "banana"}, want: false},
		{name: "Civic", args: args{str: "civic"}, want: true},
		{name: "Level", args: args{str: "level"}, want: true},
		{name: "Noon", args: args{str: "noon"}, want: true},
		{name: "Eduardo", args: args{str: "eduardo"}, want: false},
		{name: "Rodador", args: args{str: "rodador"}, want: true},
		{name: "Rir", args: args{str: "rir"}, want: true},
		{name: "Ovo", args: args{str: "ovo"}, want: true},
		{name: "15651", args: args{str: "15651"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromeFromMiddle(tt.args.str); got != tt.want {
				t.Errorf("palindromeFromMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Banana", args: args{s: "banana"}, want: "ananab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPalindromeFromMiddle_Banana(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("banana")
	}
}
func BenchmarkPalindromeReverse_Banana(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("banana")
	}
}
func BenchmarkPalindromeFromEnd_Banana(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("banana")
	}
}

func BenchmarkPalindromeFromMiddle_civic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("civic")
	}
}
func BenchmarkPalindromeReverse_civic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("civic")
	}
}
func BenchmarkPalindromeFromEnd_civic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("civic")
	}
}

func BenchmarkPalindromeFromMiddle_rodador(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("rodador")
	}
}
func BenchmarkPalindromeReverse_rodador(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("rodador")
	}
}
func BenchmarkPalindromeFromEnd_rodador(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("rodador")
	}
}
func BenchmarkPalindromeFromMiddle_nomelgibsonisacasinosbiglemon(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("nomelgibsonisacasinosbiglemon")
	}
}
func BenchmarkPalindromeReverse_nomelgibsonisacasinosbiglemon(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("nomelgibsonisacasinosbiglemon")
	}
}
func BenchmarkPalindromeFromEnd_nomelgibsonisacasinosbiglemon(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("nomelgibsonisacasinosbiglemon")
	}
}

func BenchmarkPalindromeFromMiddle_longstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("sdjfasdkflajkdslf;jaks;dlfjka;lsdjkf;lasjkdfl;ajksdlfjkalsdjkflajsdfnajsdfhquweurqjepfojiaopdsfjiaosdhufpiasdufiansdfuahsdfpuqpweufipadsnfasdhfuiaoshdufoaudfhauodfhuasdfnaiushdfuiaphfuipadnfaisudhfpadsufiahudfiauhdfiahsdfuiasuhdfpiadufnaiduhfiaduhfiahudfpiaudnciahdifphuadpfauhdpfauhpdfiaupfahusdfiphaupdifhuaipsdhufpaiudsfpihaupqiwehu9qwe8f9qwe8f9q8we9fhsdifhaidufiauhdfiahusdifuahidfuiasdufiaudfoiofduaifudsaiufdihaufidsuhaifdhuaifudiahfidshf9ew8q9f8ewq9f8ewq9uhewiqpuahipfsduiapfuhdspiauhfidpuahpifdsuhafpuaifdphuafpdhuafpdauhpfidhaicnduaipfduhaifhudaifhudianfudaipfdhusaiufdshaifdhuaifduhaifusdapfhdusiafndapiufhpaiufdhsuianfdsauhfdouahfduaofudhsoaiufhdsafnsdapifuewpqupfdshaufdsnaifudsaipfuhdsoaijfsdpoaijofpejqruewuqhfdsjanfdsjalfkjdslakjfldskja;lfdkjsal;fkjdsl;akjfld;skaj;flsdkjalfkdsafjds")
	}
}
func BenchmarkPalindromeReverse_longstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("sdjfasdkflajkdslf;jaks;dlfjka;lsdjkf;lasjkdfl;ajksdlfjkalsdjkflajsdfnajsdfhquweurqjepfojiaopdsfjiaosdhufpiasdufiansdfuahsdfpuqpweufipadsnfasdhfuiaoshdufoaudfhauodfhuasdfnaiushdfuiaphfuipadnfaisudhfpadsufiahudfiauhdfiahsdfuiasuhdfpiadufnaiduhfiaduhfiahudfpiaudnciahdifphuadpfauhdpfauhpdfiaupfahusdfiphaupdifhuaipsdhufpaiudsfpihaupqiwehu9qwe8f9qwe8f9q8we9fhsdifhaidufiauhdfiahusdifuahidfuiasdufiaudfoiofduaifudsaiufdihaufidsuhaifdhuaifudiahfidshf9ew8q9f8ewq9f8ewq9uhewiqpuahipfsduiapfuhdspiauhfidpuahpifdsuhafpuaifdphuafpdhuafpdauhpfidhaicnduaipfduhaifhudaifhudianfudaipfdhusaiufdshaifdhuaifduhaifusdapfhdusiafndapiufhpaiufdhsuianfdsauhfdouahfduaofudhsoaiufhdsafnsdapifuewpqupfdshaufdsnaifudsaipfuhdsoaijfsdpoaijofpejqruewuqhfdsjanfdsjalfkjdslakjfldskja;lfdkjsal;fkjdsl;akjfld;skaj;flsdkjalfkdsafjds")
	}
}
func BenchmarkPalindromeFromEnd_longstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("sdjfasdkflajkdslf;jaks;dlfjka;lsdjkf;lasjkdfl;ajksdlfjkalsdjkflajsdfnajsdfhquweurqjepfojiaopdsfjiaosdhufpiasdufiansdfuahsdfpuqpweufipadsnfasdhfuiaoshdufoaudfhauodfhuasdfnaiushdfuiaphfuipadnfaisudhfpadsufiahudfiauhdfiahsdfuiasuhdfpiadufnaiduhfiaduhfiahudfpiaudnciahdifphuadpfauhdpfauhpdfiaupfahusdfiphaupdifhuaipsdhufpaiudsfpihaupqiwehu9qwe8f9qwe8f9q8we9fhsdifhaidufiauhdfiahusdifuahidfuiasdufiaudfoiofduaifudsaiufdihaufidsuhaifdhuaifudiahfidshf9ew8q9f8ewq9f8ewq9uhewiqpuahipfsduiapfuhdspiauhfidpuahpifdsuhafpuaifdphuafpdhuafpdauhpfidhaicnduaipfduhaifhudaifhudianfudaipfdhusaiufdshaifdhuaifduhaifusdapfhdusiafndapiufhpaiufdhsuianfdsauhfdouahfduaofudhsoaiufhdsafnsdapifuewpqupfdshaufdsnaifudsaipfuhdsoaijfsdpoaijofpejqruewuqhfdsjanfdsjalfkjdslakjfldskja;lfdkjsal;fkjdsl;akjfld;skaj;flsdkjalfkdsafjds")
	}
}
