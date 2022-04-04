# BENCHMARK HTTP LIBS GO
Benchmark test with most popular libs in Go. Test implements a GET method wrapper to delivery a common interface.

### Libs selected:
- Fiber;
- FastHTTP (using [fastHTTP routing](http://github.com/qiangxue/fasthttp-routing), cause the package doesn't have it);
- Gin;
- Echo;
- Mux;

### Go Version:
- go version go1.18 darwin/amd64

### Setup
- CPU: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
- OS: Macos Monterey 12.3.1
- RAM: 16 GB 2667 MHz DDR4

### RESULTS
Ps.: All the benchmarks it was executed apart
<table>
    <tr>
        <th>Lib</th>
        <th>Qty Operations</th>
        <th>Average time per operation</th>
        <th>Average memory per operation</th>
    </tr>
    <tr>
        <td>Fiber</td>
        <td>2379</td>
        <td>468,03 ms/op</td>
        <td>16,42 KB/op</td>
    </tr>
    <tr>
        <td>FastHTTP routing</td>
        <td>2379</td>
        <td>489,90 ms/op</td>
        <td>16,45 KB/op</td>
    </tr>
    <tr>
        <td>Gin</td>
        <td>2730</td>
        <td>429,50 ms/op</td>
        <td>17,04 KB/op</td>
    </tr>
    <tr>
        <td>Echo</td>
        <td>2720</td>
        <td>447,75 ms/op</td>
        <td>17,02 KB/op</td>
    </tr>
    <tr>
        <td>Mux</td>
        <td>2730</td>
        <td>539,45 ms/op</td>
        <td>17,16 KB/op</td>
    </tr>
</table>
