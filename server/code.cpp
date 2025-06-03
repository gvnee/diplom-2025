#include<stdio.h>
int main(){
  int n;
  int res = 1;
  scanf("%d", &n);
  for(int i = 0;i<n;i++){
    int a;
    scanf("%d", &a);
    if(a%2) res *= a;
  }
  printf("%d\n", res);
  return 0;
}