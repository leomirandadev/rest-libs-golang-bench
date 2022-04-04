# BENCHMARK HTTP LIBS GO
Benchmark test with most popular libs in Go. Test implements a GET method wrapper to delivery a common interface.

Ps.: You can enjoy to use the HTTP dependency inversion for your clean Architecture project. 

### Libs selected:
- Fiber;
- FastHTTP (using [fastHTTP routing](http://github.com/qiangxue/fasthttp-routing), cause the package doesn't have it);
- Gin;
- Echo;
- Mux;
- Chi;

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
        <td>2347</td>
        <td>477,48 ms/op</td>
        <td>16,42 KB/op</td>
    </tr>
    <tr>
        <td>FastHTTP routing</td>
        <td>2218</td>
        <td>530,2 ms/op</td>
        <td>16,45 KB/op</td>
    </tr>
    <tr>
        <td>Gin</td>
        <td>2751</td>
        <td>409,44 ms/op</td>
        <td>15,73 KB/op</td>
    </tr>
    <tr>
        <td>Echo</td>
        <td>2660</td>
        <td>431,43 ms/op</td>
        <td>15,73 KB/op</td>
    </tr>
    <tr>
        <td>Mux</td>
        <td>2884</td>
        <td>417,61 ms/op</td>
        <td>15,73 KB/op</td>
    </tr>
    <tr>
        <td>Chi</td>
        <td>2560</td>
        <td>406,07 ms/op</td>
        <td>15,73 KB/op</td>
    </tr>
</table>
