//
// Created by darko on 01/12/2023.
//

#include "day1.h"
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int solve_day_1_a() {
    ifstream stream ("day1/input.txt");
    if (!stream) {
        cerr << "File not found." << endl;
        return 1;
    }

    string line;
    int sum = 0;
    while(getline(stream, line)) {
        erase_if(line, [](char ch) { return !isdigit(ch); } );

        sum += (line[0] - '0') * 10 + line[line.length() - 1] - '0';
    }

    stream.close();

    cout << sum << endl;
    return 0;
}
