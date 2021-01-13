package processes

import (
    "log"
    
    "../winapi"
)

func getUser(handle uintptr) (string, string) {
    thandle, err := winapi.OpenProcessToken(handle, winapi.TOKEN_QUERY)
    
    if err != nil {
        log.Fatal("processes.getUser() winapi.OpenProcessToken: ", err)
    }
    
    tokenUser, err := winapi.GetTokenInformation(thandle, winapi.TIC_TOKEN_USER)
    
    if err != nil {
        log.Fatal("processes.getUser() winapi.GetTokenInformation: ", err)
    }
    
    sid, err := winapi.ConvertSidToStringSidW(tokenUser.User.Sid)
    
    if err != nil {
        log.Fatal("processes.getUser() winapi.ConvertSidToStringSidW: ", err)
    }
    
    user, err := winapi.LookupAccountSidW("", tokenUser.User.Sid)
    
    if err != nil {
        log.Fatal("processes.getUser() winapi.LookupAccountSidW: ", err)
    }
    
    err = winapi.CloseHandle(thandle)
    
    if err != nil {
        log.Fatal("processes.getUser() winapi.CloseHandle: ", err)
    }
    
    return user, sid
}
