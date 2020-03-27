# Computer to Computer Communication

## Insecure communication
@startuml
  autonumber

  group intro
    Host1 -> Host2 : hello host2
    Host2 -> Host1 : who are you?
    Host1 -> Host2 : I am a friend of Verrol
    Host2 -> Host1 : Ok, I trust you
  end

  group exchange data
    Host1 -> Host2 : send me the secret
    Host2 -> Host1 : Here is it:"SECRET"
  end


@enduml


## Secure communication
@startuml
  autonumber

  group intro
    Host1 -> Host2 : hello host2
    Host2 -> Host1 : who are you?
    Host1 -> Host2 : I am a friend of Verrol
    Host2 -> Host1 : Prove it
    Host1 -> Host2 : something signed by Verrol
    Host2 -> Host2 : verify signature
    
    group signature valid
      Host2 -> Host1 : Ok, I trust you
      group exchange data
        Host1 -> Host2 : send me the secret
        Host2 -> Host1 : Here is it:"SECRET"
      end
  end
  
  group signature invalid
    Host2 -> Host1 : Imposter, go away
    Host1 ->x Host2
  end

@enduml