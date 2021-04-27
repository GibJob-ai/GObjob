const GET_DOGS = gql`
  query GetDogs {
    dogs {
      id
      breed
    }
  }
`;
