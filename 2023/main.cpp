#include "day1/day1.h"
#include "day2/day2.h"
#include <iostream>
#include <format>

using namespace std;

void print_delimiter() {
    cout << "+------------------------------------------------------------------------------------------+" << "\n";
}

int print_day(int day, int (*day_solution_a)(), int (*day_solution_b)()) {
    print_delimiter();
    cout << "\t Day " << day << ", part one solution:\t";
    if (day_solution_a() != 0)
        return 1;
    cout << "\t Day " << day << ", part two solution:\t";
    if (day_solution_b() != 0)
        return 1;
    print_delimiter();
    return 0;
}

int main() {
    print_delimiter();
    cout << format("{:^90}\n", "ADVENT OF CODE 2023");
    print_delimiter();

    if (print_day(1, &solve_day_1_a, &solve_day_1_b) != 0)
        return 1;

    if (print_day(2, &solve_day_2_a, &solve_day_2_b) != 0)
        return 1;

    return 0;
}
