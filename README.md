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
