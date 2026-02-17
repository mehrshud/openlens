import React, { useState, useEffect } from 'react';
import { useQuery, gql } from '@apollo/client';

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

const DASHBOARD_QUERY = gql`
  query Dashboard {
    dashboard {
      ID
      Name
      Visualizations {
        ID
        Name
        Type
      }
    }
  }
`;

const Dashboard: React.FC = () => {
  const [dashboard, setDashboard] = useState<Dashboard | null>(null);
  const { data, error, loading } = useQuery(DASHBOARD_QUERY);

  useEffect(() => {
    if (data) {
      setDashboard(data.dashboard);
    }
  }, [data]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    console.error(error);
    return <div>Error: {error.message}</div>;
  }

  if (!dashboard) {
    return <div>No dashboard data available</div>;
  }

  return (
    <div>
      <h1>{dashboard.Name}</h1>
      <ul>
        {dashboard.Visualizations.map((visualization) => (
          <li key={visualization.ID}>
            {visualization.Name} ({visualization.Type})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Dashboard;