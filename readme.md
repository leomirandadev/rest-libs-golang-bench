# BENCHMARK HTTP LIBS GO
Benchmark test with most popular libs in Go. Test implements a GET method wrapper to delivery a common interface.

### Libs selected:
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
<table>
    <tr>
        <th>Lib</th>
        <th>Qty Operations</th>
        <th>Average time per operation</th>
        <th>Average memory per operation</th>
    </tr>
    <tr>
        <td>Gin</td>
        <td>2672</td>
        <td>499074 ns/op</td>
        <td>17023 B/op</td>
    </tr>
    <tr>
        <td>Echo</td>
        <td>2762</td>
        <td>437556 ns/op</td>
        <td>17037 B/op</td>
    </tr>
    <tr>
        <td>Mux</td>
        <td>2836</td>
        <td>424428 ns/op</td>
        <td>17025 B/op</td>
    </tr>
</table>
