import React, { useState, useEffect } from 'react';
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';
import Dashboard from './components/Dashboard';

interface User {
  ID: number;
  Name: string;
  Email: string;
}

interface Config {
  Debug: boolean;
  DBURL: string;
}

interface Dashboard as IDashboard {
  ID: number;
  Name: string;
  Visualizations: Visualization[];
}

interface Visualization {
  ID: number;
  Name: string;
  Type: string;
}

const client = new ApolloClient({
  uri: 'http://localhost:8080/graphql',
  cache: new InMemoryCache(),
});

function App() {
  const [user, setUser] = useState<User | null>(null);
  const [dashboards, setDashboards] = useState<IDashboard[] | null>(null);

  useEffect(() => {
    try {
      const fetchUser = async () => {
        const response = await client.query({
          query: gql`
            query GetUser {
              user {
                ID
                Name
                Email
              }
            }
          `,
        });
        setUser(response.data.user);
      };
      fetchUser();
    } catch (error: any) {
      console.error('Error fetching user:', error.message);
    }

    try {
      const fetchDashboards = async () => {
        const response = await client.query({
          query: gql`
            query GetDashboards {
              dashboards {
                ID
                Name
                Visualizations {
                  ID
                  Name
                  Type
                }
              }
            }
          `,
        });
        setDashboards(response.data.dashboards);
      };
      fetchDashboards();
    } catch (error: any) {
      console.error('Error fetching dashboards:', error.message);
    }
  }, []);

  return (
    <ApolloProvider client={client}>
      <div>
        {user && <p>Welcome, {user.Name}!</p>}
        {dashboards && (
          <div>
            {dashboards.map((dashboard: IDashboard) => (
              <Dashboard key={dashboard.ID} dashboard={dashboard} />
            ))}
          </div>
        )}
      </div>
    </ApolloProvider>
  );
}

export default App;