import pandas as pd
from sklearn.linear_model import LogisticRegression
from sklearn.ensemble import RandomForestClassifier
import joblib

# Load dataset
data = pd.read_csv("../data/dataset.csv")
X = data.drop("target", axis=1)  # Features
y = data["target"]               # Target

# Train two models
model1 = LogisticRegression()
model1.fit(X, y)
model2 = RandomForestClassifier(n_estimators=10)
model2.fit(X, y)

# Save models
joblib.dump(model1, "models/model1.pkl")
joblib.dump(model2, "models/model2.pkl")
print("Models trained and saved!")