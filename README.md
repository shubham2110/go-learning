# Learning Go-Lang

This Repo tells actually I have learned go lang. 

I am using following links:
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals


I will be putting important takeaways in readme file for my future reference. 

## Range
we can use range for iterating over arrays/slices

```go

numbers []int

for _,number in range numbers{
    fmt.Println(number)
}
```

## Reflect.DeepEqual

We can use reflect.DeepEqual for comparing two arrays/slices.

```go

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }

```
reflect is not 'type safe' i.e. one can compare any two types

```go
if !reflect.DeepEqual("aaaa", []int{1,2,3}){

}
```

## Append

'append' function appends element behind the slice, and returns a copy of slice. 

```go
sums = append(sums, Sum(numbers))
// in above example, we has to intialize sum with new value of append(sums, Sum(numers)) because it doesn't itself appends in sums.. append function sends a copy
```

## t.Helper()

'Helper' function make sures even if error is coming in the helper function, error will be pointed out by compiter at line where helper function was called. 

```go
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
```