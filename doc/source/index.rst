Developer guide for otc-api-sign-sdk-go
============================================

.. toctree::
   :maxdepth: 10
   :hidden:

   Preparations <preparations/_index>
   Installation <installation/_index>
   Samples <samples/_index>
  

This document describes how to call cloud service APIs registered with API Gateway using the AK/SK signature authentication and Go.
You need to obtain the API information and AK/SK by referring to Obtaining API Information and :ref:`Obtaining an AK/SK  <ref_obtain_aksk>`.


- For the authentication of certain cloud service APIs that are not registered with API Gateway,
  see the :otc_docs:`API Reference <developer/api.html>` of the corresponding service.

- For the APIs provided by a cloud service, see the :otc_docs:`API Reference <developer/api.html>` of the cloud service.
  The API Reference contains a section named **Calling APIs** that describes API authentication methods.

- AK/SK authentication supports API requests with a body less than or equal to **12 MB**.
  For API requests with a larger body, **token authentication** is recommended.
  The description and example of token-based authentication are included in the section **Authentication**
  in the API reference of each cloud service.

- The local time on the client must be synchronized with the clock server to avoid a
  large offset in the value of the **X-Sdk-Date** request header.
  API Gateway checks the time format and compares the time with the time when API Gateway
  receives the request. If the time difference exceeds 15 minutes, API Gateway will reject the request.


Source Code
-----------

For source code, see :github_repo_master:`otc-api-sign-sdk-go<>` on Github.



.. note:: **Warranty Disclaimer**
    
    THE OPEN SOURCE SOFTWARE IN THIS PRODUCT IS DISTRIBUTED IN THE HOPE THAT IT
    WILL BE USEFUL,BUT WITHOUT ANY WARRANTY; WITHOUT EVEN THE IMPLIED WARRANTY
    OF MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE.

    SEE THE APPLICABLE LICENSES FOR MORE DETAILS.
