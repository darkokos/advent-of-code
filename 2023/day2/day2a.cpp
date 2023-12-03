//
// Created by darko on 02/12/2023.
//

#include "day2.h"
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <unordered_map>

using namespace std;

bool could_pull_happen(int& number, string& colour, const unordered_map<string, const int>& colour_count) {
    stringstream pull_ss(line);

    while (getline(pull_ss, line, ',')) {
        stringstream colour_ss(line);
        colour_ss >> number;
        colour_ss >> colour;
        if (colour_count.at(colour) < number)
            return false;
    }

    return true;
}

bool could_game_happen(stringstream& ss, int& number, string& colour,
                       const unordered_map<string, const int>& colour_count) {
    ss.clear();
    ss.str(line);

    while (getline(ss, line, ';')) {
        if (!could_pull_happen(number, colour, colour_count))
            return false;
    }

    return true;
}

int solve_day_2_a() {
    ifstream stream("day2/input.txt");
    if (!stream) {
        cerr << "File not found.";
        return 1;
    }

    string game;
    stringstream ss;
    int number;
    string colour;
    const unordered_map<string, const int> colour_count {
            { "red", 12 },
            { "green", 13 },
            { "blue", 14 },
    };
    int game_num;
    int sum = 0;
    while (getline(stream, line)) {
        ss.str(line);

        getline(ss, game, ':');
        getline(ss, line, ':');

        if (could_game_happen(ss, number, colour, colour_count)) {
            ss.clear();
            ss.str(game);
            ss >> line;
            ss >> game_num;
            sum += game_num;
        }

        ss.clear();
    }

    stream.close();

    cout << sum << endl;
    return 0;
}
