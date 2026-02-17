import React, { useState, useEffect } from 'react';

interface User {
  ID: number;
  Name: string;
  Email: string;
}

interface Config {
  Debug: boolean;
  DBURL: string;
}

interface Dashboard {
  ID: number;
  Name: string;
  Visualizations: Visualization[];
}

interface Visualization {
  ID: number;
  Name: string;
  Type: string;
}

function DashboardComponent(props: { dashboard: Dashboard }) {
  const [data, setData] = useState<Visualization[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('/api/graphql', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            query: `
              query {
                visualizations(dashboardId: ${props.dashboard.ID}) {
                  ID
                  Name
                  Type
                }
              }
            `,
          }),
        });

        const result = await response.json();
        setData(result.data.visualizations);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [props.dashboard.ID]);

  return (
    <div>
      <h1>{props.dashboard.Name}</h1>
      <ul>
        {data.map((visualization) => (
          <li key={visualization.ID}>{visualization.Name}</li>
        ))}
      </ul>
    </div>
  );
}

export default DashboardComponent;
