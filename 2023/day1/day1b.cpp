//
// Created by darko on 01/12/2023.
//

#include "day1.h"
#include <iostream>
#include <fstream>
#include <string>
#include <algorithm>
#include <unordered_map>

using namespace std;

const unordered_map<string, const int> digit_word_to_digit {
    { "one", 1 },
    { "two", 2 },
    { "three", 3 },
    { "four", 4 },
    { "five", 5 },
    { "six", 6 },
    { "seven", 7 },
    { "eight", 8 },
    { "nine", 9 },
};

int get_first_digit_from_string(const string& s) {
    string substr;
    for (size_t idx = 0; idx < s.length(); idx++) {
        if (isdigit(s[idx]))
            return s[idx] - '0';

        substr += s[idx];

        if (idx >= 2) {
            for (const pair<const string, const int>& kv : digit_word_to_digit) {
                if (substr.find(kv.first) != string::npos)
                    return kv.second;
            }
        }
    }
}

int get_last_digit_from_string(const string& s) {
    string substr;
    for (size_t idx = s.length() - 1; true; idx--) {
        if (isdigit(s[idx]))
            return s[idx] - '0';

        substr.insert(0, 1, s[idx]);

        if (idx <= s.length() - 3) {
            for (const pair<const string, const int>& kv : digit_word_to_digit) {
                if (substr.find(kv.first) != string::npos)
                    return kv.second;
            }
        }
    }
}

int solve_day_1_b() {
    ifstream stream ("day1/input.txt");
    if (!stream) {
        cerr << "File not found." << endl;
        return 1;
    }

    string line, line_substr;
    int sum = 0;
    while(getline(stream, line)) {
        sum += get_first_digit_from_string(line) * 10
               + get_last_digit_from_string(line);
    }

    stream.close();

    cout << sum << endl;
    return 0;
}