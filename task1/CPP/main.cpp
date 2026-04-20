#include <iostream>
#include <cmath>
#include <vector>
#include <iomanip>
#include <string>

using namespace std;

// Структура для хранения шага итерации (уточнение корня)
struct iterationStep {
    int n;
    double a;
    double b;
    double x;
    double diff; // расстояние между текущим приближением и предыдущим
};

// Функция f(x) = cos(x + 0.3) - x^2
double f(double x) {
    return cos(x + 0.3) - pow(x, 2);
}

// Производная f'(x) = -sin(x + 0.3) - 2x
double df(double x) {
    return -sin(x + 0.3) - 2 * x;
}

// Функция для метода простых итераций x = phi(x)
// x = sqrt(cos(x + 0.3))
double phi(double x) {
    return sqrt(cos(x + 0.3));
}

void printBorder(int width) {
    cout << "+" << setfill('-') << setw(width) << "" << "+" << setfill(' ') << endl;
}

void printTable(const string& title, const vector<iterationStep>& history) {
    int tableWidth = 55; // Общая ширина внутреннего пространства
    
    cout << "\n=== " << title << " ===\n";

    printBorder(tableWidth);

    cout << "| " << left << setw(6) << "N" 
         << setw(16) << "a_n (x_n)" 
         << setw(16) << "b_n (x_n+1)" 
         << setw(17) << "diff" << "|" << endl;

    cout << "| " << setfill('-') << setw(tableWidth - 2) << "" << setfill(' ') << "  |" << endl;

    cout << fixed << setprecision(6);
    for (const auto& step : history) {
        cout << "| " << left << setw(6) << step.n 
             << setw(16) << step.a 
             << setw(16) << step.b 
             << setw(17) << step.diff << "|" << endl;
    }

    printBorder(tableWidth);
}

int main() {
    double a = 0.0, b = 1.0;
    double eps = 1e-4;

    // Метод половинного деления
    vector<iterationStep> bisection;
    double currA = a, currB = b; // текущие значения
    int k = 0;
    while ((currB - currA) > eps) {
        double mid = (currA + currB) / 2.0;
        bisection.push_back({++k, currA, currB, mid, currB - currA});
        if (f(currA) * f(mid) < 0) currB = mid;
        else currA = mid;
    }
    printTable("Метод половинного деления", bisection);

    // Метод Ньютона
    vector <iterationStep> newton;
    double xPrev = 0.5; // Начальное приближение
    double xNext;
    k = 0;
    // do-while цикл с постусловием
    do {
        xNext = xPrev - f(xPrev) / df(xPrev);
        newton.push_back({++k, xPrev, xNext, xNext, abs(xNext - xPrev)});
        if (abs(xNext - xPrev) < eps) break;
        xPrev = xNext;
    } while (true);
    printTable("Метод Ньютона", newton);

    // Метод простых итераций
    vector<iterationStep> fixedPoint;
    xPrev = 0.5;
    k = 0;
    do {
        xNext = phi(xPrev);
        fixedPoint.push_back({++k, xPrev, xNext, xNext, abs(xNext - xPrev)});
        if (abs(xNext - xPrev) < eps) break;
        xPrev = xNext;
    } while (true);
    printTable("Метод простых итераций", fixedPoint);
    return 0;
}
