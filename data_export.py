
import pandas as pd
import json

def export_data(data, format):
    if format == 'csv':
        pd.DataFrame(data).to_csv('export.csv', index=False)
    elif format == 'json':
        with open('export.json', 'w') as f:
            json.dump(data, f, indent=4)
    else:
        raise ValueError('Unsupported export format')

def main():
    # Example usage:
    data = {'name': ['John', 'Alice', 'Bob'], 'age': [25, 30, 35]}
    export_data(data, 'csv')
    export_data(data, 'json')

if __name__ == '__main__':
    main()
