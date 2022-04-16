1. Client iniitates request to the Core gRPC server. Currently allows only getParam, getParams, putParam, putParams.
2. Core initialtes Auth validation to the AuthEngine.
    - AuthEngine checks if token exists in RedisCache.
        - if it does, returns ok to core.
        - if it doesn't initiates connection to OrmEngine with a query for information.
            - If it's in the OrmEngine and we got a response with details, put the details in redis.
            - If it's not in the OrmEngine, assumes authentication failure. (need some prevention here)
    - Assuming everything is okay, respond with ok.
3. Core initiates connection with PolEngine, provinging it with all the request metadata and userid
    - PolEngine goes through every metadata detail and checks for all maching policies configured for the org.
    - Should I build one massive query to the OrmEngine or create many small ones?
    - Leave place for an optional redis policy cache here. (invalidate on SetPol)
4. Core initiates request to ParamEngine.
    1. Send only the policies and OrgId will fetch all params configured for those policies under an organization. (getParamValuesByPolicy) - DB Does the work
    2. Send policies and provide a parameter will fetch the parameter fron Orm and then match the right value based on the provided policies (getParamValueByName) - Pod does the work.

