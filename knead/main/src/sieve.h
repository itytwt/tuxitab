#pragma once
#include "skeleton.h"

/* @snippet:sieve */
namespace snippet {

	using namespace std;

	inline vector< int > sieve(const ll n) {
		vector< bool > is_prime(uint(n), true);
		vector< int > prime;

		is_prime[0] = is_prime[1] = false;
		for (ll i = 2; i < n; ++i) {
			if (is_prime[uint(i)])
				prime.phb(int(i));

			for (ll j = 0; i * prime[uint(j)] < n; ++j) {
				is_prime[uint(i * prime[uint(j)])] = false;
				if (i % prime[uint(j)] == 0)
					break;
			}
		}

		return prime;
	}

}
/* @endsnippet */