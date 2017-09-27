int calibrationTime = 30;
long unsigned int lowIn;
long unsigned int pause = 500; 

boolean lockLow = true;
boolean takeLowTime; 
 
int pirPin = 7;

void setup(){
  Serial.begin(9600);
  pinMode(pirPin, INPUT);
  pinMode(LED_BUILTIN, OUTPUT);
  digitalWrite(pirPin, LOW);
  for(int i = 0; i < calibrationTime; i++){
    delay(1000);
  }
  Serial.println("SENSOR ACTIVE");
  delay(50);
}

void loop() {
  if(digitalRead(pirPin) == HIGH){
      digitalWrite(LED_BUILTIN, HIGH);
      delay(100);
      Serial.println("BANG!");
  }
  else {
      digitalWrite(LED_BUILTIN, LOW);
      delay(100);   
  }
}