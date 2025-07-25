# 责任链模式

在日常生活中，我们经常会申请各种各样的权限。每个审批过程中又会出现多个节点，每个节点都会有对应的同学进行审批。每个权限节点的同学通常只负责审批自己权限内的资源。

来看一个带有一些逻辑的责任链：报销审批流程。

公司为了更高效、安全规范地把控审核工作，通常会将整个审批工作过程按负责人或者工作职责进行拆分，并组织好各个环节中的逻辑关系及走向，最终形成标准化的审批流程。审批流程需要依次通过财务专员、财务经理、财务总监的审批。如果申请金额在审批人的审批职权范围内则审批通过并终止流程，反之则会升级至更高层级的上级去继续审批，直至最终的财务总监，如果仍旧超出财务总监的审批金额则驳回申请，流程终止。但是实际的业务场景中这套审批流程可能会由于业务的变动，而不断变化。话说要方便地对业务链条进行拆分、重组，以及对单独节点的增、删、改。结构松散的业务处理节点让系统具备更加灵活的可伸缩性、可扩展性。

责任链模式（Chain of Responsibility Pattern）用于将请求从一个处理程序传递到另一个处理程序，直到找到能够处理该请求的处理程序为止。在 Go 语言中，责任链模式通常通过函数闭包和链式调用实现，提高代码的灵活性和可扩展性

>在示例中，创建一个简单的审批流程链，包括经理、总监和 CFO 三个层级的审批。每个层级有不同的审批金额上限，根据请求金额决定是否批准请求或将请求传递给下一级审批人。
