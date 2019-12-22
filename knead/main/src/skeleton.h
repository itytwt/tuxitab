#pragma once
/* @snippet:skeleton */
/* @order:0 */
#include <iostream>
#include <algorithm>
#include <string>
#include <vector>
#include <cmath>
#include <map>
#ifdef debug
#define ifdebug(x) {x}
#else
#define ifdebug(x)
#endif
#define watch(x) std::cerr << (#x) << " = " << (x) << std::endl
#define all(x) (x).begin(), (x).end()
#define mkp make_pair
#define fr first
#define sc second
#define phb(x) push_back(x)

template<typename T>
constexpr auto sz(T x) { return (x).size(); }

typedef unsigned int uint;
typedef long long ll;
typedef std::pair< int, int > pii;
const double kPi = 2 * std::acos(0.0);
/* @endsnippet */