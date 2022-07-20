#ifndef __ERRORS_H__
#define __ERRORS_H__

#define MAX_ERR_NAME_LENGTH 30
#define SOLVE_ERR 1
#define MUTEX_LOCK_ERR 2
char errors[][MAX_ERR_NAME_LENGTH] = {
    "SUCCESS.\n",
    "command solve failed.\n",
    "lock special mutex failed.\n"
};

#endif