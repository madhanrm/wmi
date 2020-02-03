﻿// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.wmi.Common
{
    public abstract class WmiReference : IWmiBase
    {
        public WmiReference(string reference)
        {
            Reference = reference;
        }
        public string Reference { get; set; }
    }
}
