### Interesting moments: 
- I tried to get from `https://www.rusprofile.ru/ajax.php?&query=${TaxId}&action=search`, but
  output json doesn't have one attribute `KPP`. So my decision is to parse `https://www.rusprofile.ru/search?query=${TaxId}`.
- _custom jsonUnmarshaller_ is way better in case of main json struct has slice inside.
  It would be easier to get json attributes in slice. 
  [Article](https://jhall.io/posts/go-json-tricks-array-as-structs/) with good examples.

- 