## Recursive Example

Clone this repo
And run command ```go test -v ./...```

Result:
```bash
=== RUN   TestRecite
=== RUN   TestRecite/Read_and_print_lyrics
=== RUN   TestRecite/Recite_by_1
This is Twinkle, twinkle, little star
=== RUN   TestRecite/Recite_by_2
This is How I wonder what you are and Twinkle, twinkle, little star
=== RUN   TestRecite/Recite_by_4
This is Like a diamond in the sky and Up above the world so high and How I wonder what you are and Twinkle, twinkle, little star
=== RUN   TestRecite/Recite_by_empty_param_/_random
--- PASS: TestRecite (0.00s)
    --- PASS: TestRecite/Read_and_print_lyrics (0.00s)
    --- PASS: TestRecite/Recite_by_1 (0.00s)
    --- PASS: TestRecite/Recite_by_2 (0.00s)
    --- PASS: TestRecite/Recite_by_4 (0.00s)
    --- PASS: TestRecite/Recite_by_empty_param_/_random (0.00s)
PASS
ok      github.com/aginanjar/simple-tdd-recursive/src   0.171s
```