#pragma once
#include "skeleton.h"

/* @snippet:math */
namespace snippet {

	using namespace std;

	inline ll inv(ll a, ll m) {
		ll t0 = 0, t = 1;
		ll r0 = m, r = a;
		while (r != 0) {
			ll q = r0 / r, s;
			s = t0, t0 = t, t = s - q * t;
			s = r0, r0 = r, r = s - q * r;
		}
		return (t0 % m + m) % m;
	}

}
/* @endsnippet */