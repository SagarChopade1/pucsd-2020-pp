#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include "./operation/addition.c"
#include "./operation/subtraction.c"
#include "./operation/multiplication.c"
#include "./operation/division.c"
#include "./operation/mod.c"
#include "value-stack.c"
#include "operator-stack.c"
#define MAX_SIZE 100
#define BUF 11
char precidence[]={'/','*','+','-'};

void endofList()
{
  while(opetop!=NULL)
  {
       float var;
       char array[10],topelement=opetop->operator;

       if(topelement=='-'){
        var=subtration(atof(valpop()),atof(valpop()));

       }   
       else if(topelement=='+')
       {
        var =addition(atof(valpop()),atof(valpop()));
       }   
       else if(topelement=='*')
        {
          var=multi(atof(valpop()),atof(valpop()));
        } 
       else if(topelement=='/')
           var=division(atof(valpop()),atof(valpop()));
       else
       {
          operatorpop();
       }
   
    if(('('!=topelement) && (')'!=topelement))
    {
    sprintf(array, "%f", var);
    valpush(array);
    opetop=opetop->nextoperator;
    }  
             
  }
}

// void middleexp()
// {
    
// }
void token(char * line)
{
    //printf("Inside token %s\n",line);
    int charindex=0,flag=1;
    char string[BUF];
    while(flag)
    {
        int asciival=(int)line[charindex];
        //printf("new line is catch %c %d\n",line[charindex],(int)line[charindex]);
        if(((asciival>=40 && asciival<=57) || asciival==37) && asciival!=44 /*","*/)
        {
          if((asciival>=40 &&  asciival<=43) || (asciival>=45 &&  asciival<=47) || asciival==37 )
          { 
            
            operatorpush(line[charindex]);
            //printf("%c\n",line[charindex]);
          }
          else 
          {
              //concatination
              char tempvar[2];
              tempvar[0]=line[charindex];
              tempvar[1]='\0';
              strcat(string,tempvar);
              //printf("%s\n",string);   
          }
        }
        else if(asciival==32 || asciival==36)
        {
            //printf("(%s) string to posh is \n",string);
            if(strlen(string)!=0)
               valpush(string);
            //printf("Stirng os :%s \n",string);
            string[0]='\0';

            if(asciival==36)
            { 
                flag=0;
                endofList();
            }
        }
        else
        {   if(asciival!=0){
             printf("Invalide input (%d)(%c)in input string %s\n",(int)line[charindex],line[charindex],line);
             break;  
            }          
        }
        charindex++;
    }

}
void printtwostack()
{
  printf("values stack is :\n");
  valdisplay(valtop);
  printf("operator is : \n");
  oprdisplay(opetop);
}
int main(int argv,char * argc)
{
    char str[MAX_SIZE];
    printf("Enter string to calculate\n");
    scanf("%[^\n]s",str);
    str[(int)strlen(str)]='$';
    //str[(int)strlen(str)]=("a");
    //printf("string is :(%s)(%d)",str,(int)strlen(str));
    token(str);
    printf("result is equal to : %s \n",valtop->value);
    //printtwostack();
    // printf("addition is %f\n",addition(10.0,5.0));
    // printf("subtraction is %f\n",subtration(10.0,5.0));
    // printf("mult is %f\n",multi(10.0,5.0));
    // printf("divi is %f\n",division(10.0,5.0));
    // printf("mod is %f\n",mod(10,5));
    return 0;
}
