import React from 'react';
import { Visualization } from '../../models';
import { useQuery, gql } from '@apollo/client';

interface VisualizationComponentProps {
  visualization: Visualization;
}

const GET_VISUALIZATION_DATA = gql`
  query GetVisualizationData($id: Int!) {
    visualization(id: $id) {
      id
      name
      type
    }
  }
`;

const VisualizationComponent: React.FC<VisualizationComponentProps> = ({ visualization }) => {
  const { data, error, loading } = useQuery(GET_VISUALIZATION_DATA, {
    variables: { id: visualization.ID },
  });

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    console.error(error);
    return <div>Error fetching visualization data</div>;
  }

  if (!data) {
    return <div>No data available for visualization</div>;
  }

  const visualizationData = data.visualization;

  return (
    <div>
      <h2>{visualizationData.name}</h2>
      <p>Type: {visualizationData.type}</p>
    </div>
  );
};

export default VisualizationComponent;