# Benchmarks

A list of Benchmarks that I'm curious about in terms of performance

## Float to String

Convert a `float64` into a string with different precisions(b,e,E,f,g,G,x,X)

```
goos: linux
goarch: amd64
pkg: github.com/xescugc/benchmarks/floattostring
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkSprintf-8       5636593               208.0 ns/op            16 B/op          2 allocs/op
BenchmarkSprintf-8       5894581               216.8 ns/op            16 B/op          2 allocs/op
BenchmarkSprintf-8       6012067               209.9 ns/op            16 B/op          2 allocs/op
BenchmarkSprintf-8       5492665               204.8 ns/op            16 B/op          2 allocs/op
BenchmarkSprintf-8       6029721               199.4 ns/op            16 B/op          2 allocs/op
BenchmarkStrconv_b-8    24572958                67.49 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_b-8    11599827                94.14 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_b-8    15319095                74.81 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_b-8    17128976                69.93 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_b-8    16832437                68.96 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_e-8    16952355                71.33 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_e-8    17252976                68.69 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_e-8    17677388                68.87 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_e-8    16997462                68.60 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_e-8    17035735                67.28 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_E-8    16982416                78.06 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_E-8    12102391                88.08 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_E-8    15318040                76.21 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_E-8    13588233                82.40 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_E-8    15589762                71.88 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_f-8     5560080               200.2 ns/op            24 B/op          1 allocs/op
BenchmarkStrconv_f-8     6141081               194.5 ns/op            24 B/op          1 allocs/op
BenchmarkStrconv_f-8     6213342               193.1 ns/op            24 B/op          1 allocs/op
BenchmarkStrconv_f-8     6224721               194.0 ns/op            24 B/op          1 allocs/op
BenchmarkStrconv_f-8     4919584               218.4 ns/op            24 B/op          1 allocs/op
BenchmarkStrconv_g-8    18818922                66.66 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_g-8    17550420                66.85 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_g-8    18010244                70.68 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_g-8    16070859                69.92 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_g-8    16830454                66.49 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_G-8    18819770                67.81 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_G-8    16773321                80.71 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_G-8    15486588                70.52 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_G-8    17720232                66.51 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_G-8    17345986                65.80 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_x-8    25846693                44.40 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_x-8    26201896                43.73 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_x-8    25843628                43.74 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_x-8    26508726                46.10 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_x-8    26057958                45.82 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_X-8    27176410                44.62 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_X-8    25395074                45.41 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_X-8    28679840                45.11 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_X-8    26853253                43.85 ns/op           24 B/op          1 allocs/op
BenchmarkStrconv_X-8    26579210                47.34 ns/op           24 B/op          1 allocs/op
```

## Int to string

Convert 3 integers int string and concatenate them together

```
goos: linux
goarch: amd64
pkg: github.com/xescugc/benchmarks/inttostring
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkStrconvItoaConcat-8            28900945                36.93 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaConcat-8            31847365                38.34 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaConcat-8            31966470                36.30 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaConcat-8            30419577                36.75 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaConcat-8            31234812                36.35 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaSprintf-8           32385788                36.19 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaSprintf-8           27632836                41.71 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaSprintf-8           19554723                57.54 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaSprintf-8           20835058                54.16 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvItoaSprintf-8           21874614                53.86 ns/op            8 B/op          1 allocs/op
BenchmarkStrconvParseIntConcat-8         5409883               223.9 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntConcat-8         5044908               225.0 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntConcat-8         5410629               217.6 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntConcat-8         5453209               221.4 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntConcat-8         4926210               229.1 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntSprintf-8        4967634               223.4 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntSprintf-8        5400036               224.0 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntSprintf-8        5365278               219.9 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntSprintf-8        5407722               231.5 ns/op            56 B/op          4 allocs/op
BenchmarkStrconvParseIntSprintf-8        5455396               221.9 ns/op            56 B/op          4 allocs/op
BenchmarkSprintf-8                       8396949               138.6 ns/op             8 B/op          1 allocs/op
BenchmarkSprintf-8                       8697976               172.4 ns/op             8 B/op          1 allocs/op
BenchmarkSprintf-8                       8029159               140.8 ns/op             8 B/op          1 allocs/op
BenchmarkSprintf-8                       8956730               132.6 ns/op             8 B/op          1 allocs/op
BenchmarkSprintf-8                       8744660               132.7 ns/op             8 B/op          1 allocs/op
```
