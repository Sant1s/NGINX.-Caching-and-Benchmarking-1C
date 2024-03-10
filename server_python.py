from flask import Flask, jsonify, request
from datetime import datetime

app = Flask(__name__)


@app.route('/api/date')
def get_date():
    current_date = datetime.now().strftime("%Y-%m-%d")
    data = [{"year": current_date.split("-")[0], "month": current_date.split("-")[1],
             "day": current_date.split("-")[2]}] * 10000
    return jsonify(data)


@app.route('/api/name', methods=['POST'])
def get_name():
    name = request.form.get('name')
    data = [{"name": name}] * 10000
    return jsonify(data)


if __name__ == '__main__':
    app.run(debug=True)
