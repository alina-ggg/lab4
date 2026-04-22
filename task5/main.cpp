#include <iostream>
#include <vector>
#include <random>

using namespace std;

typedef bool (*strategyFunc)(int, const vector<bool>&, const vector<bool>&);

// 1. Случайный выбор
bool randomChoice(int round_number, const vector<bool>& self_choices, const vector<bool>& enemy_choices) {
    // Используем статический генератор, чтобы не инициализировать его каждый раунд
    static mt19937 gen(random_device{}());
    static uniform_int_distribution<> dist(0, 1);
    return dist(gen) == 1; // 1 - сотрудничать (true), 0 - предать (false)
}

// 2. Взаимные ходы
bool titForTat(int round_number, const vector<bool>& self_choices, const vector<bool>& enemy_choices) {
    if (round_number == 0) return true;
    return enemy_choices.back();
}

// 3. Если его один раз предали, то предает всегда
bool grimTrigger(int round_number, const vector<bool>& self_choices, const vector<bool>& enemy_choices) {
    // Используем Range-based for по требованию лабы
    for (bool choice : enemy_choices) {
        if (!choice) return false;
    }
    return true;
}

void runSession(strategyFunc s1, strategyFunc s2, const char* name1, const char* name2) {
    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<> dist(100, 200);
    int totalRounds = dist(gen);

    vector<bool> history1, history2;
    int score1 = 0, score2 = 0;

    for (int i = 0; i < totalRounds; ++i) {
        bool move1 = s1(i, history1, history2);
        bool move2 = s2(i, history2, history1);

        if (move1 && move2) {
            score1 += 24; score2 += 24;
        } else if (!move1 && !move2) {
            score1 += 4; score2 += 4;
        } else if (!move1 && move2) {
            score1 += 20; score2 += 0;
        } else {
            score1 += 0; score2 += 20;
        }

        history1.push_back(move1);
        history2.push_back(move2);
    }

    cout << name1 << " vs " << name2 << ": " << score1 << " - " << score2 << " (" << totalRounds << " раундов)" << endl;
}

int main() {
    cout << "Результаты 3 запусков (со случайной стратегией):" << endl;

    // 1 vs 2
    // Очки будут разными, так как 2 будет отвечать на каждое предательство
    runSession(randomChoice, titForTat, "randomChoice", "titForTat");

    // 2 vs 3
    // очки могут остаться одинаковыми, если оба начнут с мира и никто не "спровоцирует"
    runSession(titForTat, grimTrigger, "titForTat", "grimTrigger");

    // 3 vs 1
    // будет самая большая разница: как только случайный предаст, 3 уйдет в агрессию до конца
    runSession(grimTrigger, randomChoice, "grimTrigger", "randomChoice");

    return 0;
}