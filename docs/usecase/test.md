
```uml
@startuml register_user
participant Browser
box "motting" #LightBlue
participant mottingAPIServer as api
participant WebPushAPIServer as push
participant mottingDataBase as apidb
participant WebPushDataBase as pushdb
participant WebPushBatch as batch
end box
participant GooglePushServer

Browser -> api: RegisterUserRequest
api -> apidb: RegisterUser
apidb --> api: Result
api --> Browser: RegisterUserResponse
@enduml
```
