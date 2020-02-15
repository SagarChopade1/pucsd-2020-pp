struct dataval
{
    char *value;
    struct dataval *nextval;
};
typedef struct dataval dataval;
dataval *valtop;
void initialize()
{
    valtop = NULL ;
} 

char *copyString(char *str)
{
  char *tmp = malloc(strlen(str) + 1);
  if (tmp)
    strcpy(tmp, str);
  return tmp;
}

void valpush(char *val)
{
  dataval *temp;
  temp=malloc(sizeof(dataval));
  temp->value=copyString(val);
  temp->nextval=valtop;
  valtop=temp;

}
char * valpop()
{
    dataval *temp;
    temp=valtop;
    valtop=temp->nextval;
    char *retval=malloc(sizeof(char)*strlen(temp->value));
    strcpy(retval,temp->value);
    free(temp->value);
    free(temp);
    return retval;
}

void valdisplay(dataval *head)
{
    if(head == NULL)
    {
        printf("NULL\n");
    }
    else
    {
        printf("%s\n", head -> value);
        valdisplay(head->nextval);
    }
}