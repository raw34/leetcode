//一、01背包问题
// 二维
#include<iostream>
#include<algorithm>

using namespace std;

const int N = 1010;

int n, m;
int f[N][N], w[N], v[N];

int main()
{
    cin >> n >> m;
    for(int i = 1; i <= n; i++) cin >> w[i] >> v[i];

    for(int i = 1; i <= n; i++)
        for(int j = 1; j <= m; j++)
        {
            f[i][j] = f[i-1][j];
            if(j >= w[i])
                f[i][j] = max(f[i][j], f[i-1][j - w[i]] + v[i]);
        }

    cout << f[n][m] << endl;
    return 0;
}

// 一维
#include <iostream>
#include <algorithm>

using namespace std;

const int N = 1010;

int n, m;
int v[N], w[N];
int f[N];

int main()
{
    cin >> n >> m;

    for (int i = 1; i <= n; i ++ ) cin >> v[i] >> w[i];

    for (int i = 1; i <= n; i ++ )
        for (int j = m; j >= v[i]; j -- )
            f[j] = max(f[j], f[j - v[i]] + w[i]);

    cout << f[m] << endl;

    return 0;
}

//二、完全背包问题
// 二维暴力
#include<iostream>

using namespace std;

const int N=1010;

int f[N][N];
int w[N],v[N];

int main()
{
    int n, m;
    cin >> n >> m;
    for(int i = 1; i <= n; i++)  cin >> w[i] >> v[i];

    for(int i = 1; i <= n; i++)
        for(int j = 1; j <= m; j++)
            for(int k = 0; k*w[i] <= j; k++)
                f[i][j] = max(f[i][j], f[i-1][j - k*w[i]] + k*v[i]);

    cout << f[n][m] << endl;
    return 0;
}

// 二维
#include<iostream>
#include<algorithm>

using namespace std;

const int N=1010;


int f[N][N];
int n, m;
int w[N],v[N];

int main()
{
    cin >> n >> m;
    for(int i = 1; i <= n; i++) cin >> w[i] >> v[i];

    for(int i = 1; i <= n; i++)
        for(int j = 1; j <= m; j++)
        {
            f[i][j] = f[i-1][j];
            if(j >= w[i])
                f[i][j]=max(f[i][j], f[i][j - w[i]] + v[i]);
        }

    cout<< f[n][m] <<endl;
    return 0;
}

// 一维
#include <iostream>
#include <algorithm>

using namespace std;

const int N = 1010;

int n, m;
int v[N], w[N];
int f[N];

int main()
{
    cin >> n >> m;
    for (int i = 1; i <= n; i ++ ) cin >> v[i] >> w[i];

    for (int i = 1; i <= n; i ++ )
        for (int j = v[i]; j <= m; j ++ )
            f[j] = max(f[j], f[j - v[i]] + w[i]);

    cout << f[m] << endl;

    return 0;
}

//三、多重背包问题I
// 二维
#include<iostream>
#include<algorithm>

using namespace std;

const int N=110;

int f[N][N], s[N], w[N], v[N];

int main()
{
    int n, m;
    cin >> n >> m;
    for(int i = 1; i <= n; i++) cin >> w[i] >> v[i] >> s[i];

    for(int i = 1; i <= n; i++)
        for(int j = 1; j <= m; j++)
            for(int k = 0; k <= s[i] && k*w[i] <= j;k++)
                f[i][j]=max(f[i][j], f[i-1][j - k*w[i]] + k*v[i]);

    cout<< f[n][m] <<endl;
    return 0;
}

// 一维
#include <iostream>
#include <algorithm>

using namespace std;

const int N = 110;

int n, m;
int v[N], w[N], s[N];
int f[N][N];

int main()
{
    cin >> n >> m;

    for (int i = 1; i <= n; i ++ ) cin >> v[i] >> w[i] >> s[i];

    for (int i = 1; i <= n; i ++ )
        for (int j = 0; j <= m; j ++ )
            for (int k = 0; k <= s[i] && k * v[i] <= j; k ++ )
                f[i][j] = max(f[i][j], f[i - 1][j - v[i] * k] + w[i] * k);

    cout << f[n][m] << endl;
    return 0;
}

//四、多重背包问题II
#include <iostream>
#include <algorithm>

using namespace std;

const int N = 11010, M = 2010;

int n, m;
int v[N], w[N];
int f[N];

int main()
{
    cin >> n >> m;

    int cnt = 0;
    for (int i = 1; i <= n; i ++ )
    {
        int a, b, s;
        cin >> a >> b >> s;
        int k = 1;
        while (k <= s)
        {
            cnt ++ ;
            v[cnt] = a * k;
            w[cnt] = b * k;
            s -= k;
            k *= 2;
        }
        if (s > 0)
        {
            cnt ++ ;
            v[cnt] = a * s;
            w[cnt] = b * s;
        }
    }

    n = cnt;
    for (int i = 1; i <= n; i ++ )
        for (int j = m; j >= v[i]; j -- )
            f[j] = max(f[j], f[j - v[i]] + w[i]);

    cout << f[m] << endl;

    return 0;
}

//五、分组背包问题
// 二维
#include<bits/stdc++.h>
using namespace std;

const int N=110;
int f[N][N];  //只从前i组物品中选，当前体积小于等于j的最大值
int v[N][N], w[N][N], s[N];   //v为体积，w为价值，s代表第i组物品的个数
int n, m, k;

int main(){
    cin >> n >> m;
    for(int i = 1; i <= n; i++){
        cin >> s[i];
        for(int j = 0; j < s[i]; j++){
            cin >> v[i][j] >> w[i][j];  //读入
        }
    }

    for(int i = 1; i <= n; i++){
        for(int j = 0; j <= m; j++){
            f[i][j] = f[i-1][j];  //不选
            for(int k = 0; k < s[i]; k++){
                if(j >= v[i][k])
                   f[i][j] = max(f[i][j], f[i-1][j -v[i][k]] + w[i][k]);
            }
        }
    }
    cout << f[n][m] << endl;
}

// 一维
#include <iostream>
#include <algorithm>

using namespace std;

const int N = 110;

int n, m;
int v[N][N], w[N][N], s[N];
int f[N];

int main()
{
    cin >> n >> m;

    for (int i = 1; i <= n; i ++ )
    {
        cin >> s[i];
        for (int j = 0; j < s[i]; j ++ )
            cin >> v[i][j] >> w[i][j];
    }

    for (int i = 1; i <= n; i ++ )
        for (int j = m; j >= 0; j -- )
            for (int k = 0; k < s[i]; k ++ )
                if (v[i][k] <= j)
                    f[j] = max(f[j], f[j - v[i][k]] + w[i][k]);

    cout << f[m] << endl;

    return 0;
}