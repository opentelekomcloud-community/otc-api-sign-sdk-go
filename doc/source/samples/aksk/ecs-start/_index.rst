Start ECS using AK/SK signing
=================================

.. toctree::
   :maxdepth: 3
   :includehidden:


Sample on how to start an ECS instance using an ak/sk request signing.


For complete source code, see :github_repo_master:`samples-doc/aksksigning/ecs-start<samples-doc/aksksigning/ecs-start>` on GitHub.

Requirements
------------

.. list-table:: Environment variables
    :widths: 20 20 25
    :header-rows: 1

    * - Environment variable name
      - Value
      - Remarks

    * - ECS_INSTANCE_ID
      - <ID of ecs instance>
      - ID of the ecs instance to start

    * - OTC_SDK_PROJECT_ID
      - <Project ID>
      - Needed if ecs instance is in a sub project see :api_usage:`Obtaining a Project ID<guidelines/calling_apis/obtaining_required_information.html#obtaining-a-project-id>`

    * - OTC_SDK_AK
      - <Access Key>
      - see: :api_usage:`Generating AK and SK<guidelines/calling_apis/ak_sk_authentication/generating_an_ak_and_sk.html#apig-en-api-180328005>`

    * - OTC_SDK_SK
      - <Secret Key>
      - see: :api_usage:`Generating AK and SK<guidelines/calling_apis/ak_sk_authentication/generating_an_ak_and_sk.html#apig-en-api-180328005>`

Installation and Running
-------------------------

.. code-block:: bash
   :caption: Install and run the sample

   # Running the sample
   go run .

