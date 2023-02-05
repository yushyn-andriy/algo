#include<bits/stdc++.h>
using namespace std;


int  solve(string & s){
    vector<stack<char> >vs;

    for(int i=0;i<s.size();i++){
        if(vs.empty()){
            stack<char>st;
            vs.push_back(st);
            vs[0].push(s[i]);
            continue;
        }
    bool flag=false;
   /* 
     flag will be used to check whether we found any container whose top 
     is greater than the current container
    */   
        int mini=-1;
      // mini will be used to store the index of nearest appropriate container stack
        for(int j=0;j<vs.size();j++){
            if(s[i]<=vs[j].top()){
                flag=true;
                if(mini==-1)
                    mini=j;
                else if(vs[mini].top()>vs[j].top()){
                    mini=j;
                }
            }
        }
if(!flag)
        {
            stack<char>st;
            vs.push_back(st);
            vs[vs.size()-1].push(s[i]);
        }
else{  //if we could not able to place the container on any of previous container
            vs[mini].push(s[i]);
        }
    }
return vs.size();//returning size of vector of stacks
}

int main()
{
    string s;
    cin>>s;
    int i=1;
    while(s!="end"){
        cout<<"Case "<<i++<<": "<<solve(s)<<endl;
cin>>s;
    }
}