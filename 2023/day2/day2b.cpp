//
// Created by darko on 02/12/2023.
//

#include "day2.h"
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <algorithm>
#include <chrono>

using namespace std;
using namespace chrono;

int get_set_power(stringstream& ss, string& colour, string& number, int& max_r, int& max_g, int& max_b) {
    max_r = 0;
    max_g = 0;
    max_b = 0;

    while (getline(ss, number, ' ')) {
        getline(ss, colour, ' ');
        if (colour[colour.length() - 1] == ';' || colour[colour.length() - 1] == ',')
            colour = colour.substr(0, colour.length() - 1);

        if (colour == "red") {
            max_r = max(max_r, stoi(number));
        }
        else if (colour == "green") {
            max_g = max(max_g, stoi(number));
        }
        else {
            max_b = max(max_b, stoi(number));
        }
    }

    return max_r * max_g * max_b;
}

int solve_day_2_b() {
    auto start = high_resolution_clock::now();

    ifstream stream("day2/input.txt");
    if (!stream) {
        cerr << "File not found.";
        return 1;
    }

    stringstream ss;
    string colour, number;
    int max_r, max_g, max_b, sum = 0;
    while (getline(stream, line)) {
        ss.str(line);

        getline(ss, line, ':');
        getline(ss, line, ':');

        ss.clear();
        ss.str(line.substr(1, line.length() - 1));

        sum += get_set_power(ss, colour, number, max_r, max_g, max_b);

        ss.clear();
    }

    stream.close();

    cout << sum << endl;
    return 0;
}
