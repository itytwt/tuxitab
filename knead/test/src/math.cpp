#include "../pch.h"
using namespace std;

TEST(TestMath, TestInv) {
	ll a = 3, m = 11;
	ll expected = 4;
	ll actual = snippet::inv(a, m);
	EXPECT_EQ(expected, actual);

	a = 10, m = 17;
	expected = 12;
	actual = snippet::inv(a, m);
	EXPECT_EQ(expected, actual);
}