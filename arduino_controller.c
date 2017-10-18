int calibrationTime = 30;
long unsigned int lowIn;
long unsigned int pause = 500; 

boolean lockLow = true;
boolean takeLowTime; 
 
int pirPin = 4;

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
  if(digitalRead(pirPin) == LOW){
      delay(100);
  }
  else {
      Serial.println("BANG!");
      delay(100);
  }
}