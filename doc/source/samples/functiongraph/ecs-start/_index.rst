Start ECS using FunctionGraph and agency
========================================

.. toctree::
   :maxdepth: 3
   :includehidden:

Sample on how to start an ECS instance using an agency.

This sample show how to use functiongraph start an ECS instance using 
following credentials retrieved by the agency:


* ctx.GetSecurityAccessKey()
* ctx.GetSecuritySecretKey()
* ctx.GetSecurityToken()


For complete source code, see :github_repo_master:`samples-doc/functiongraph/fg-ecs-start<samples-doc/functiongraph/fg-ecs-start>` on GitHub.

Deployment
----------

Build deployment zip
^^^^^^^^^^^^^^^^^^^^^^

To build the deployment zip, execute following command in
folder: **samples-doc/functiongraph/fg-ecs-stop**

.. code-block:: bash

   make all


This will create **deploy.zip** with code and dependencies included.

Create FunctionGraph function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Create function
*******************

Use `OpentelekomCloud FunctionGraph console <https://console.otc.t-systems.com/functiongraph/>`_ to create a function with following
settings:


**Create With**:  Create from scratch 

**Basic Information**

* **Function Type**  Event Function  
* **Region**  <YOUR REGION>  
* **Function Name** <YOUR FUNCTION NAME>  
* **Agency**  Specify an agency with policy to start ECS instance  
* **Runtime**  Go 1.x

Upload code
*******************

Use **Upload** -> **Local ZIP** and upload *deploy.zip* from previous step.

Configure function
*******************

In **Configuration** -> **Basic Settings** -> **Handler**:   *startECS*

In **Configuration** -> **Environment Variables** add following variables:

.. list-table:: Environment variables
    :widths: 20 20 25
    :header-rows: 1

    * - Environment variable name
      - Value
      - Remarks

    * - INSTANCE_ID
      - <ID of ecs instance>
      - ID of ECS instance to start

    * - ECS_ENDPOINT
      - <ecs endpoint>
      - Default: ecs.eu-de.otc.t-systems.com,
        see :otc_docs:`Regions and Endpoints<regions-and-endpoints/index.html>`


Create Test Event
*******************

In **Code** create a Test Event using "Blank Template" (Event is not used in function).

Test function
^^^^^^^^^^^^^^^

Click **Test** to test function.

