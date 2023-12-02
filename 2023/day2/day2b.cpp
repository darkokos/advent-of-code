//
// Created by darko on 02/12/2023.
//

#include "day2.h"
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int solve_day_2_b() {
    ifstream stream("day2/input.txt");
    if (!stream) {
        cerr << "File not found.";
        return 1;
    }

    string line;
    int sum = 0;
    while (getline(stream, line)) {

    }

    stream.close();

    cout << sum << endl;
    return 0;
}
