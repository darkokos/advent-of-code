//
// Created by darko on 02/12/2023.
//

#include "day2.h"
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
#include <unordered_map>

using namespace std;

const unordered_map<string, int> colour_count {
        { "red", 12 },
        { "green", 13 },
        { "blue", 14 },
};
stringstream ss;
string line;
int number;
string colour;

bool could_pull_happen() {
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

bool could_game_happen() {
    ss.clear();
    ss.str(line);

    while (getline(ss, line, ';')) {
        if (!could_pull_happen())
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
    int game_num;
    int sum = 0;
    while (getline(stream, line)) {
        ss.str(line);

        getline(ss, game, ':');
        getline(ss, line, ':');

        if (could_game_happen()) {
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
