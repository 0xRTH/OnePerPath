# OnePerPath

A simple command-line tool that deduplicates URLs based on their domain, path, and extension. It's particularly useful when you want to keep only one representative URL per unique path pattern.

## Features

- Deduplicates URLs based on domain + directory path + extension
- Treats numeric path segments (like dates) as equivalent
- Preserves the first encountered URL for each unique pattern

## Usage

```bash
# Process URLs from a file
cat urls.txt | ./OnePerPath

# Or pipe URLs directly
echo -e "https://example.com/2020/11/file.jpg\nhttps://example.com/2021/12/other.jpg" | ./OnePerPath
```

## Example

Given these input URLs:
```
https://example.com/uploads/2020/11/image1.jpg
https://example.com/uploads/2021/12/image2.jpg
https://other.com/uploads/2020/11/image1.jpg
```

Only these will be output (as they have unique domain+path+extension combinations):
```
https://example.com/uploads/2020/11/image1.jpg
https://other.com/uploads/2020/11/image1.jpg
```

Note: All numeric path segments (like years and months) are treated as equivalent. 