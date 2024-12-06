#include <iostream>
#include <fstream>
#include <map>
#include <vector>
#include <cmath>

using std::ifstream;
using std::map;
using std::vector;
using std::abs;
using std::cout;
using std::sort;
using std::stoi;
using std::string;
using std::getline;

int main() {
    ifstream inFile;
    // inFile.open("2024/day01/example.txt");
    inFile.open("2024/day01/input.txt");

    if (!inFile.is_open()) {
        cout << "Unable to open file";
        return 1;
    }

    string line;
    string delimiter = " ";
    vector<int> left, right;
    map<int, int> leftM, rightM;
    while (getline(inFile, line)) {
        int leftInt = stoi(line.substr(0, line.find(delimiter)));
        int rightInt = stoi(line.substr(line.find_last_of(delimiter) + 1, line.npos));
        left.push_back(leftInt);
        right.push_back(rightInt);
        if (leftM.find(leftInt) == leftM.end()) {
            leftM[leftInt] = 0;
        }
        leftM[leftInt] += 1;
        if (rightM.find(rightInt) == rightM.end()) {
            rightM[rightInt] = 0;
        }
        rightM[rightInt] += 1;
    }

    sort(left.begin(), left.end());
    sort(right.begin(), right.end());

    int sum = 0;
    for (int i = 0; i < left.size(); i++) {
        sum += abs(left[i] - right[i]);
    }

    cout << sum <<'\n';

    sum = 0;
    int def;
    map<int, int>::iterator leftI = leftM.begin();
    while (leftI != leftM.end()) {
        if (rightM.find(leftI->first) != rightM.end()) {
            def = rightM.find(leftI->first)->second;
        }
        sum += leftI->first * leftI->second * def;
        def = 0;
        leftI++;
    }

    cout << sum << '\n';

    inFile.close();

    return 0;
}
