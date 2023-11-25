type bluerpc = {
  users: {
    bye: {
      query: (queryParams: { query: string }) => void;
      mutation: (input: {
        queryParams: { query: string };
        input: { house: string };
      }) => {
        fieldOneOut: string;
        fieldTwoOut: string | undefined;
        fieldThreeOut: string;
      };
    };
    depth1: {
      bye: {
        query: (queryParams: { query: string }) => void;
        mutation: (input: {
          queryParams: { query: string };
          input: { house: string };
        }) => {
          fieldOneOut: string;
          fieldTwoOut: string | undefined;
          fieldThreeOut: string;
        };
      };
      depth2: {
        bye: {
          query: (queryParams: { query: string }) => void;
          mutation: (input: {
            queryParams: { query: string };
            input: { house: string };
          }) => {
            fieldOneOut: string;
            fieldTwoOut: string | undefined;
            fieldThreeOut: string;
          };
        };
        depth3: {
          bye: {
            query: (queryParams: { query: string }) => void;
            mutation: (input: {
              queryParams: { query: string };
              input: { house: string };
            }) => {
              fieldOneOut: string;
              fieldTwoOut: string | undefined;
              fieldThreeOut: string;
            };
          };
          hello: {
            query: (queryParams: { query: string }) => void;
            mutation: (input: {
              queryParams: { query: string };
              input: { house: string };
            }) => {
              fieldOneOut: string;
              fieldTwoOut: string | undefined;
              fieldThreeOut: string;
            };
          };
        };
        hello: {
          query: (queryParams: { query: string }) => void;
          mutation: (input: {
            queryParams: { query: string };
            input: { house: string };
          }) => {
            fieldOneOut: string;
            fieldTwoOut: string | undefined;
            fieldThreeOut: string;
          };
        };
      };
      hello: {
        query: (queryParams: { query: string }) => void;
        mutation: (input: {
          queryParams: { query: string };
          input: { house: string };
        }) => {
          fieldOneOut: string;
          fieldTwoOut: string | undefined;
          fieldThreeOut: string;
        };
      };
    };
    hello: {
      query: (queryParams: { query: string }) => void;
      mutation: (input: {
        queryParams: { query: string };
        input: { house: string };
      }) => {
        fieldOneOut: string;
        fieldTwoOut: string | undefined;
        fieldThreeOut: string;
      };
    };
  };
};
