
import matplotlib.pyplot as plt
import pandas as pd

class DataVisualization:
    def __init__(self, data):
        self.data = data

    def plot_bar_chart(self, x, y):
        plt.bar(self.data[x], self.data[y])
        plt.xlabel(x)
        plt.ylabel(y)
        plt.show()

    def plot_line_chart(self, x, y):
        plt.plot(self.data[x], self.data[y])
        plt.xlabel(x)
        plt.ylabel(y)
        plt.show()

    def plot_scatter_plot(self, x, y):
        plt.scatter(self.data[x], self.data[y])
        plt.xlabel(x)
        plt.ylabel(y)
        plt.show()
