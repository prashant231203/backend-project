import sys
import json
import joblib

# Load models
model1 = joblib.load("models/model1.pkl")
model2 = joblib.load("models/model2.pkl")

# Get input from command line (passed by Go)
input_data = json.loads(sys.argv[1])  # Expecting a list like [1.0, 2.0, 3.0]

# Predict
pred1 = model1.predict([input_data])[0]
pred2 = model2.predict([input_data])[0]

# Output as JSON
result = {"model1": int(pred1), "model2": int(pred2)}
print(json.dumps(result))