
# OpenLens
OpenLens is a data analysis tool.

## Features
- Data analysis
- Data visualization

## Usage
To use the data visualization module, create an instance of the DataVisualization class and call the desired plot function.
python
from data_visualization import DataVisualization
import pandas as pd

data = pd.DataFrame({
    'x': [1, 2, 3],
    'y': [4, 5, 6]
})

visualization = DataVisualization(data)
visualization.plot_bar_chart('x', 'y')
