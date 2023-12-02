cmake_minimum_required(VERSION 3.26)
project(2023)

set(CMAKE_CXX_STANDARD 23)

add_executable(2023 main.cpp
        day1/day1.h
        day1/day1a.cpp
        day1/day1b.cpp
        day2/day2.h
        day2/day2a.cpp
        day2/day2b.cpp
)
