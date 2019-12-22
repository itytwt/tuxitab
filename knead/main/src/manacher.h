#pragma once
#include "skeleton.h"

/* @snippet:manacher */
namespace snippet {

	using namespace std;

	inline vector< int > manacher(const string &s) {
		string t = "^";
		int c, r;

		for (int i = 0; i < int(s.size()); ++i)
			t += "#", t += s[i];
		t += "#$";

		vector< int > f(2 * int(s.size()) + 3, 0);
		f[0] = c = r = 0;
		for (int i = 1; i < int(t.size()) - 1; ++i) {
			int mr = c - (i - c);

			f[i] = (r >= i ? min(f[mr], r - i) : 0);
			while (t[i + (f[i] + 1)] == t[i - (f[i] + 1)])
				++f[i];

			if (i + f[i] > r)
				c = i, r = i + f[i];
		}

		return f;
	}

}
/* @endsnippet */