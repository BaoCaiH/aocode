#include <iostream>
#include <fstream>
#include <map>
#include <vector>
#include <cmath>
#include <tuple>

using std::abs;
using std::cout;
using std::getline;
using std::ifstream;
using std::map;
using std::sort;
using std::stoi;
using std::string;
using std::tie;
using std::tuple;
using std::vector;

vector<string> split(const string &s, const string &delimeter = " ")
{
    vector<string> tokens;
    string token;
    size_t start = 0, end, delimeter_length = delimeter.length();
    while ((end = s.find(delimeter, start)) != string::npos)
    {
        token = s.substr(start, end - start);
        tokens.push_back(token);
        start = end + delimeter_length;
    }
    tokens.push_back(s.substr(start));

    return tokens;
}

vector<int> parseLine(vector<string> tokens)
{
    vector<int> res;
    for (string s : tokens)
    {
        res.push_back(stoi(s));
    }

    return res;
}

tuple<bool, int> validate(vector<int> level)
{
    int size = level.size();
    bool increasing = false;
    for (int i = 0; i < size; i++)
    {
        if (i == 0 && level[i] < level[i + 1])
        {
            increasing = true;
        }
        if (i == size - 1)
        {
            return {true, -1};
        }
        if (increasing && level[i + 1] - level[i] <= 3 && level[i + 1] > level[i])
        {
            continue;
        }
        if (!increasing && level[i] - level[i + 1] <= 3 && level[i + 1] < level[i])
        {
            continue;
        }
        return {false, i};
    }
    return {false, -1};
}

template<typename T>
vector<T> slice(vector<T> const &v, int m, int n)
{
    auto first = v.cbegin() + m;
    auto last = v.cbegin() + n + 1;
 
    return vector<T>(first, last);
    // return vec;
}

template<typename T>
vector<T> copyExcept(vector<T> const &v, int except)
{
    vector<T> vec;
    for (int i = 0; i < v.size(); i++)
    {
        if (i != except) {vec.push_back(v[i]);}
    }
    return vec;
}

int main()
{
    ifstream inFile;
    // inFile.open("2024/day02/example.txt");
    inFile.open("2024/day02/input.txt");

    if (!inFile.is_open())
    {
        cout << "Unable to open file";
        return 1;
    }

    string line;
    string delimiter = " ";
    int sum = 0, idx;
    bool res, res0, res1, res2;
    vector<int> level;
    while (getline(inFile, line))
    {
        level = parseLine(split(line));
        tie(res, idx) = validate(level);
        int throwaway;
        if (res)
        {
            sum += 1;
        }
        else // Part 2
        {
            tie(res0, throwaway) = validate(copyExcept(level, idx - 1));
            tie(res1, throwaway) = validate(copyExcept(level, idx));
            tie(res2, throwaway) = validate(copyExcept(level, idx + 1));
            if (res0 || res1 || res2) {sum += 1;}
        }
    }

    cout << sum << '\n';

    inFile.close();

    return 0;
}
