
# Assignment for Day 1

Parse command line arguments to extract the first name, last name, middle name (if provided), and the country code.
Determine the name format based on the country code. Reorder the name according to the determined format. Output the reordered name.


## About Name Formats

Only 7 countries in the world use the format `family name, given name` (Eastern name order) as the standard name format.

- [Reddit](https://www.reddit.com/r/MapPorn/comments/786ori/name_order_in_countries_4500x2234/)
- [Wikipedia](https://en.wikipedia.org/wiki/Personal_name#:~:text=The%20order%20family%20name%2C%20given,Laos%2C%20Indonesia%2C%20Malaysia%20or%20the)


## How to run

```bash
go build

./exercise-01 Xuan Van Hong VN
```

## Expected output

```bash
Output: Van Hong Xuan
```

### Test

```bash
go test
```

