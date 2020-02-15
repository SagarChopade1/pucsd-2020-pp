struct dataoperator
{
    char operator;
    struct dataoperator *nextoperator;
};
typedef struct dataoperator dataoperator;
dataoperator *opetop;

void oprinitialize()
{
    opetop = NULL ;
} 
void operatorpush(char oper)
{
    dataoperator *temp;
    temp=malloc(sizeof(dataoperator));
    temp->operator=oper;
    temp->nextoperator=opetop;
    opetop=temp;
}

char operatorpop()
{
    dataoperator *temp;
    temp=opetop;
    opetop=temp->nextoperator;
    char retval;
    retval=temp->operator;
    //free(temp->nextoperator);
    free(temp);
    return retval;
}

void oprdisplay(dataoperator *head)
{
    if(head == NULL)
    {
        printf("NULL\n");
    }
    else
    {
        printf("%c\n", head->operator);
        oprdisplay(head->nextoperator);
    }
}